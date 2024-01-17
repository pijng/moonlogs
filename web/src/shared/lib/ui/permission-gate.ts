import { userModel } from "@/entities/user";
import { UserRole } from "@/shared/api/users";
import { combine, createStore } from "effector";
import { variant } from "forest";

export const PermissionGate = (role: UserRole, fn: () => void) => {
  const $access = combine(
    createStore(role),
    userModel.$currentAccount,
    (requiredRole, { role }): { state: "allowed" | "restricted" } => {
      return { state: requiredRole === role ? "allowed" : "restricted" };
    },
  );

  variant({
    source: $access,
    key: "state",
    cases: {
      allowed: () => fn(),
      restricted: () => null,
    },
  });
};
