export type Input = {
  username: string
  latitude: number
  longitude: number
}

const API = 'http://localhost:8081'

export const createLocation = async (input: Input) => {
  const url = `${API}/location`
  await fetch(url, {
    method: 'POST',
    body: JSON.stringify(input)
  })
}