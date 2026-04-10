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

    route.push(`/map?username=${username}`)
  }

  return (
    <div className="flex flex-col flex-1 items-center justify-center bg-zinc-50 font-sans dark:bg-black">

      <Card className="w-[500px] p-10">
        <CardHeader>
          <CardTitle>
            <span className="font-bold">Olá, seja bem vindo(a)</span>
          </CardTitle>
          <CardDescription>Entre com o nome de usuário</CardDescription>
        </CardHeader>
        <CardContent>
          <form 
            className="flex flex-col gap-2"
            onSubmit={e => handleOnSubmit(e)}>
            <div className="flex flex-col gap-2">
              <Label>
                <span  className="font-normal">
                  Nome de usuário
                </span>
              </Label>
              <Input value={username} onChange={e => setUsername(e.target.value)} />
            </div>
            <Button type="submit" className="cursor-pointer">Entrar!</Button>
          </form>
        </CardContent>
      </Card>
    </div>
  );
}
