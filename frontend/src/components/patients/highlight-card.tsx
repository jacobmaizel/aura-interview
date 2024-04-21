"use client";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import { Input } from "../ui/input";
import { useState } from "react";
import { Button } from "../ui/button";
import { Patient } from "@/lib/types";
import { formatDate } from "@/lib/utils";

type HighlightCardProps = {
  patient: Patient;
};

export function HighlightCard({ patient }: HighlightCardProps) {
  const [temperature, setTemperature] = useState<number>(0);
  const [error, setError] = useState<string>("");
  const [message, setMessage] = useState<string>("");

  const handleSubmit = async () => {
    if (temperature === 0) {
      setError("Temperature is required");
      return;
    }

    if (temperature > 110) {
      setError("Temperature must be less than 110");
      return;
    }
    setError("");
    await fetch(`http://localhost:5050/api/v1/temperatures/`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        temperature,
        patient_id: patient.id,
        date: formatDate(new Date()),
      }),
    })
      .then((res) => {
        if (res.ok) {
          setTemperature(0);
          setError("");
          setMessage("Temperature submitted successfully");
        } else {
          setError("You can only submit one temperature per day per client");
        }
      })
      .catch((error) => {
        console.error("Error:", error);
        setError("You can only submit one temperature per day per client");
      });
  };
  return (
    <Card>
      <CardHeader>
        <CardTitle>Highlights</CardTitle>
        <CardDescription>Take a temperature reading</CardDescription>
      </CardHeader>
      <CardContent>
        <Input
          placeholder="Temperature"
          type="number"
          value={temperature}
          min={0}
          max={110}
          onChange={(e) => setTemperature(parseFloat(e.target.value))}
        />

        <Button className="mt-4" onClick={handleSubmit}>
          Submit
        </Button>

        {message && <p className="text-green-500 mt-2">{message}</p>}

        {error && <p className="text-red-500 mt-2">{error}</p>}
      </CardContent>
    </Card>
  );
}
