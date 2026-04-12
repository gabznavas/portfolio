'use client'

import { FeatureCollection, Point } from 'geojson'
import { Position } from '@/models/location';
import { createLocation } from '@/service/location/create-location/create-location';
import { listLocationByPosition } from '@/service/location/list-location-by-position/list-location-by-position';
import maplibregl from 'maplibre-gl'
import { useSearchParams } from 'next/navigation';
import { useEffect, useRef, useState } from 'react'

const FIVE_SECONDS = 5000

export default function MapPage() {
  const searchParams = useSearchParams()

  const mapRef = useRef<HTMLDivElement | null>(null);
  const [map, setMap] = useState<maplibregl.Map | null>(null)

  const username = searchParams.get('username')

  const [positions, setPositions] = useState<Position[]>([])

  useEffect(() => {
    setInterval(() => {
      navigator.geolocation.getCurrentPosition(
        async position => {
          const lat = position.coords.latitude
          const long = position.coords.longitude
          if (!username) return

          console.log(username, lat, long)
          await createLocation({
            username,
            latitude: lat,
            longitude: long
          })

          const locations = await listLocationByPosition({
            lat,
            long,
            rangeKm: 500
          })

          setPositions(prev => {
            const newPositions = [...prev]
            for (let location of locations) {
              const foundIndex = newPositions.findIndex(item => item.username.trim() === location.username.trim())
              if (foundIndex >= 0) {
                newPositions[foundIndex].latitude = location.latitude
                newPositions[foundIndex].longitude = location.longitude
              } else {
                newPositions.unshift(location)
              }
            }

            console.log(newPositions)
            return newPositions
          })
        },
        err => {
          debugger
          console.error(err)
        }
      )
    }, FIVE_SECONDS)
  }, [])

  useEffect(() => {
    if (!mapRef || !mapRef.current) return;
    const map = new maplibregl.Map({
      container: mapRef.current,
      style: 'https://tiles.openfreemap.org/styles/liberty',
      center: [-51.22, -22.22],
      zoom: 13,
    })
    setMap(map)

    return () => {
      setMap(null)
      map.remove()
    };
  }, [])

  useEffect(() => {
  if (!map) return

  if (!map.isStyleLoaded()) {
    map.on('load', () => updateMap())
    return
  }

  updateMap()

  function updateMap() {
    if(!map) return
    const geojson: FeatureCollection<Point> = {
      type: 'FeatureCollection',
      features: positions.map(position => ({
        type: 'Feature',
        properties: {
          username: position.username
        },
        geometry: {
          type: 'Point',
          coordinates: [position.longitude, position.latitude],
        }
      }))
    }

    if (map.getSource('point')) {
      const source = map.getSource('point') as maplibregl.GeoJSONSource
      source.setData(geojson)
      return
    }

    map.addSource('point', {
      type: 'geojson',
      data: geojson
    })

    map.addLayer({
      id: 'point',
      type: 'circle',
      source: 'point',
      paint: {
        'circle-radius': 10,
        'circle-color': '#3887be'
      }
    })

    map.addLayer({
      id: 'point-username',
      type: 'symbol',
      source: 'point',
      layout: {
        'text-field': ['get', 'username'],
        'text-size': 12,
        'text-offset': [0, 1.5],
        'text-anchor': 'top'
      },
      paint: {
        'text-color': '#000000'
      }
    })
  }
}, [positions, map])

  return (
    <div ref={mapRef} style={{ height: '100vh' }} />
  )
}