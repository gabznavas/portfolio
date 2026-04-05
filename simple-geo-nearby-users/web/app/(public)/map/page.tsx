'use client'

import maplibregl from 'maplibre-gl'
import { useEffect, useRef } from 'react'

export default function MapPage() {
  const mapRef = useRef<HTMLDivElement | null>(null);

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