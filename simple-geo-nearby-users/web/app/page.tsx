'use client'

import { Button } from "@/components/ui/button";
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useRouter } from "next/navigation";
import { SubmitEvent, useState } from "react";

export default function Home() {
  const route = useRouter()

  const [username, setUsername] = useState('')

  const handleOnSubmit = (e: SubmitEvent<HTMLFormElement>) => {
    e.preventDefault()

    if (!username) {
      alert('nome de usuário é necessário');
      return;
    }

    if (username.length > 20) {
      alert('Nome de usuário deve ser menor que 20 caracteres.')
      return;
    }

    route.push('/map')
  }

  return (
    <div className="flex flex-col flex-1 items-center justify-center bg-zinc-50 font-sans dark:bg-black">

      <Card>
        <CardHeader>
          <CardTitle>Olá, seja bem vindo(a)</CardTitle>
          <CardDescription>Entre com o nome de usuário</CardDescription>
        </CardHeader>
        <CardContent>
          <form onSubmit={e => handleOnSubmit(e)}>
            <div>
              <Label>Nome de usuário</Label>
              <Input value={username} onChange={e => setUsername(e.target.value)} />
            </div>

            <Button type="submit">Entrar!</Button>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
