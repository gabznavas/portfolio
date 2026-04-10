import { API } from "@/service/utils"

export type Input = {
  username: string
  latitude: number
  longitude: number
}


export const createLocation = async (input: Input) => {
  const url = `${API}/location`
  await fetch(url, {
    method: 'POST',
    body: JSON.stringify(input)
  })
}