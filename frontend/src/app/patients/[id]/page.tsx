import { Card, CardContent, CardHeader, CardTitle } from "@/components/ui/card";
import { Medication, Patient } from "@/lib/types";

import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { HighlightCard } from "@/components/patients/highlight-card";
import { TemperatureChartCard } from "@/components/patients/temperature-chart-card";

type PageProps = {
  params: {
    id: string;
  };
};
export default async function Page({ params }: PageProps) {
  const patient: Patient = await fetch(
    `http://backend:5050/api/v1/patients/${params.id}`
  ).then((res) => res.json());

  const medications: Medication[] = await fetch(
    `http://backend:5050/api/v1/medications/patient/${params.id}`
  ).then((res) => res.json());

  return (
    <main className="flex min-h-screen flex-col p-24 gap-4">
      <Card>
        <CardHeader>
          <CardTitle>Patient Information</CardTitle>
        </CardHeader>
        <CardContent>
          <p>Name: {patient.name}</p>
          <p>First Name: {patient.first_name}</p>
          <p>Age: {patient.age}</p>
          <p>Height: {patient.height}</p>
          <p>Weight: {patient.weight}</p>
        </CardContent>
      </Card>

      <div className="grid grid-cols-1 gap-4 md:grid-cols-2">
        <HighlightCard patient={patient} />

        <Card>
          <CardHeader>
            <CardTitle>Temperature Chart</CardTitle>
          </CardHeader>
          <CardContent>
            <TemperatureChartCard patient={patient} />
          </CardContent>
        </Card>
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Medications</CardTitle>
        </CardHeader>

        <CardContent>
          {medications?.length > 0 ? (
            <>
              <Table>
                <TableHeader>
                  <TableRow>
                    <TableHead className="">Name</TableHead>
                    <TableHead>Dosage</TableHead>
                    <TableHead>Start Date</TableHead>
                    <TableHead>End Date</TableHead>
                  </TableRow>
                </TableHeader>
                <TableBody>
                  {medications?.map((medication) => (
                    <TableRow
                      className="cursor-pointer"
                      // onClick={() => router.push(`/patients/${patient.id}`)}
                      key={medication.id}
                    >
                      <TableCell>{medication.name}</TableCell>
                      <TableCell>{medication.dosage}</TableCell>
                      <TableCell>{medication.start_date}</TableCell>
                      <TableCell>{medication.end_date}</TableCell>
                    </TableRow>
                  ))}
                </TableBody>
              </Table>
            </>
          ) : (
            <p>No medications found</p>
          )}
        </CardContent>
      </Card>
    </main>
  );
}
