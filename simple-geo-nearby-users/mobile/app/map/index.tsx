import { Camera, MapView } from "@maplibre/maplibre-react-native";

export default function MapIndex() {
  return (
    <MapView 
      mapStyle='https://tiles.openfreemap.org/styles/liberty'
      style={{ flex: 1 }}>
      <Camera
        zoomLevel={14}
        centerCoordinate={[-51.22, -22.22]}
      />
    </MapView>
  )
}