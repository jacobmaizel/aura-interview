import { PatientTable } from "@/components/patients/table";
import { Patient } from "@/lib/types";
import Image from "next/image";

export default async function Home() {
  const patients: Patient[] = await fetch(
    "http://backend:5050/api/v1/patients"
  ).then((res) => res.json());
  return (
    <main className="flex min-h-screen flex-col p-24 ">
      <div>
        <PatientTable patients={patients} />
      </div>
    </main>
  );
}
