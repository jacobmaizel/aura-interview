import { Patient } from "@/lib/types";

type PatientDetailProps = {
  patient: Patient;
};
export function PatientDetail({ patient }: PatientDetailProps) {
  return <div>{patient.name}</div>;
}
