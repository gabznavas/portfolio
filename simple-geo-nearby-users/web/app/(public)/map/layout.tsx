import 'maplibre-gl/dist/maplibre-gl.css';

type Props = {
  children: React.ReactNode
}

export default function MapLayout(props: Props) {
  return (
    <div>
      {props.children}
    </div>
  )
}