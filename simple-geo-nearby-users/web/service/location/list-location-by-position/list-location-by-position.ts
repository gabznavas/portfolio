import { Position } from "@/models/location"
import { API } from "@/service/utils"

type Input = {
  lat: number
  long: number
  rangeKm: number
}

type Output = Position[]

export const listLocationByPosition = async (input: Input): Promise<Output> => {
  const url = `${API}/location?latitude=${input.lat}&longitude=${input.long}&radiusKm=${input.rangeKm}`
  const response = await fetch(url)
  const data = await response.json()
  return data.locations
}
