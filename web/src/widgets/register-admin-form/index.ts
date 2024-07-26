import { RegisterAdmin } from "@/features/register-admin";
import { Header } from "@/shared/ui";

export const RegisterAdminForm = () => {
  Header("Create initial moonlogs admin");
  RegisterAdmin();
};
