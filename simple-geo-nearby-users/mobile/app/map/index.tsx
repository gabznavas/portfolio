import { useLocalSearchParams } from 'expo-router';
import { Camera, CircleLayer, MapView, ShapeSource } from "@maplibre/maplibre-react-native";
import { useEffect, useState } from 'react';
import { Position } from '@/models/location';
import { FeatureCollection, Point } from 'geojson';
import * as Location from 'expo-location';
import { createLocation } from '@/service/location/create-location/create-location';
import { listLocationByPosition } from '@/service/location/list-location-by-position/list-location-by-position';

export default function MapIndex() {
  const { username } = useLocalSearchParams();

  const [position, setPosition] = useState<Position[]>([]);

  useEffect(() => {
    let interval: number;

    (async () => {
      const { status } = await Location.requestForegroundPermissionsAsync();

      if (status !== 'granted') {
        console.log('Permissão negada');
        return;
      }
    })()

    interval = setInterval(async () => {
      console.log('interval')
      if (!username) return
      const { status } = await Location.requestForegroundPermissionsAsync();

      if (status !== 'granted') {
        console.log('Permissão negada');
        return;
      }

      // pegar localização atual
      const location = await Location.getCurrentPositionAsync({});

      await createLocation({
        latitude: location.coords.latitude,
        longitude: location.coords.longitude,
        username: username as string,
      })

      const locations = await listLocationByPosition({
        lat: location.coords.latitude,
        long: location.coords.longitude,
        rangeKm: 500
      })

      console.log(locations)

      setPosition((prev) => {
        const map = new Map<string, Position>();

        // antigos
        prev.forEach((p) => map.set(p.username, p));

        // novos (sobrescreve)
        locations.forEach((p) => map.set(p.username, p));

        return Array.from(map.values());
      });
    }, 5000)

    return () => {
      if (interval) clearInterval(interval);
    }
  }, []);

  return (
    <MapView
      mapStyle="https://tiles.openfreemap.org/styles/liberty"
      style={{ flex: 1 }}
    >
      <Camera
        // zoomLevel={14}
        centerCoordinate={[-51.22, -22.22]}
      />

      <ShapeSource id="pointSource" shape={{
        type: 'FeatureCollection',
        features: position.map((p) => ({
          type: 'Feature',
          geometry: {
            type: 'Point',
            coordinates: [p.longitude, p.latitude],
          },
          properties: {
            username: p.username,
          },
        })),
      }}>
        <CircleLayer
          id="pointLayer"
          style={{
            circleRadius: 8,
            circleColor: '#ff0000',
          }}
        />
      </ShapeSource>
    </MapView>
  );
}