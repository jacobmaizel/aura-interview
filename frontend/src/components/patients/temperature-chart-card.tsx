"use client";
import { BodyTemperature, Patient } from "@/lib/types";
import { useEffect, useState } from "react";
import {
  Legend,
  Line,
  LineChart,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from "recharts";
import {
  Select,
  SelectContent,
  SelectItem,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { formatDate, formatShortDate } from "@/lib/utils";
type TemperatureChartCardProps = {
  patient: Patient;
};
type ScaleOptions = "1" | "3" | "6";
export function TemperatureChartCard({ patient }: TemperatureChartCardProps) {
  const [monthScale, setMonthScale] = useState<ScaleOptions>("1");
  const [data, setData] = useState<BodyTemperature[]>([]);

  useEffect(() => {
    fetch(
      `http://localhost:5050/api/v1/temperatures/patient/${patient.id}` +
        `?scale=${monthScale}`,
      {
        method: "GET",
        headers: {
          "Content-Type": "application/json",
        },
      }
    )
      .then((res) => res.json())
      .then((data) => {
        const formattedData = data.map((item: BodyTemperature) => {
          return {
            ...item,
            date: formatShortDate(item.date),
          };
        });
        setData(
          formattedData.sort((a: any, b: any) => a.date.localeCompare(b.date))
        );
      })
      .catch((error) => {
        console.error("Error:", error);
      });
  }, [monthScale, patient]);

  return (
    <div>
      <Select onValueChange={(val) => setMonthScale(val as ScaleOptions)}>
        <SelectTrigger className="w-[180px]">
          <SelectValue placeholder="Month Scale" />
        </SelectTrigger>
        <SelectContent>
          <SelectItem value="1">1</SelectItem>
          <SelectItem value="3">3</SelectItem>
          <SelectItem value="6">6</SelectItem>
        </SelectContent>
      </Select>
      <ResponsiveContainer width="100%" height={300}>
        <LineChart
          margin={{ top: 20, right: 0, left: 0, bottom: 0 }}
          data={data}
        >
          <XAxis dataKey="date" />
          <YAxis />
          <Tooltip />
          <Legend />
          <Line name={"Temperatures"} dataKey="temperature" />
        </LineChart>
      </ResponsiveContainer>
    </div>
  );
}
