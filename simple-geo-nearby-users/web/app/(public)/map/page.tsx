'use client'

import { createLocation } from '@/service/location/create-location/create-location';
import { listLocationByPosition } from '@/service/location/list-location-by-position/list-location-by-position';
import maplibregl from 'maplibre-gl'
import { useSearchParams } from 'next/navigation';
import { useEffect, useRef } from 'react'

const FIVE_SECONDS = 5000

export default function MapPage() {
  const searchParams = useSearchParams()

  const mapRef = useRef<HTMLDivElement | null>(null);

  const username = searchParams.get('username')

  useEffect(() => {
    setInterval(() => {
      navigator.geolocation.getCurrentPosition(
        async position => {
          const lat = position.coords.latitude   
          const long = position.coords.longitude
          if(!username) return
          
          console.log(username, lat, long)
          await createLocation({
            username,
            latitude: lat,
            longitude : long
          })

          const locations = await listLocationByPosition({
            lat,
            long,
            rangeKm: 50
          })
          console.log('locations', locations)
        },
        err => {
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

    return () => map.remove();
  }, [])

  return (
    <div ref={mapRef} style={{ height: '100vh' }} />
  )
}