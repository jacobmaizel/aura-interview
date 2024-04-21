"use client";
import { Patient } from "@/lib/types";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { useRouter } from "next/navigation";
type TableProps = {
  patients: Patient[];
};

// name: string;
// first_name: string;
// age: number;
// height: number;
// weight: number;
// gender: string;
export function PatientTable({ patients }: TableProps) {
  const router = useRouter();
  return (
    <div>
      <h1 className="text-3xl font-bold">Patient List</h1>
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead className="w-[100px]">Name</TableHead>
            <TableHead>First Name</TableHead>
            <TableHead>Age</TableHead>
            <TableHead>Height</TableHead>
            <TableHead>Weight</TableHead>
            <TableHead>Gender</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {patients.map((patient) => (
            <TableRow
              className="cursor-pointer"
              onClick={() => router.push(`/patients/${patient.id}`)}
              key={patient.id}
            >
              <TableCell>{patient.name}</TableCell>
              <TableCell>{patient.first_name}</TableCell>
              <TableCell>{patient.age}</TableCell>
              <TableCell>{patient.height}</TableCell>
              <TableCell>{patient.weight}</TableCell>
              <TableCell>{patient.gender}</TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </div>
  );
}
