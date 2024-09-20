export interface Item {
  id: string;
  title: string;
  price: number;
  quantity: number;
  status: "PENDING" | "APPROVED" | "REJECTED";
  owner_id: string;
}