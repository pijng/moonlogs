import { RegisterAdmin } from "@/features";
import { Header } from "@/shared/ui";

export const RegisterAdminForm = () => {
  Header("Create initial moonlogs admin");
  RegisterAdmin();
};
