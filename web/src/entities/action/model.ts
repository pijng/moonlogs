import { Action, getAction, getActions } from "@/shared/api";
import { createEffect, createStore } from "effector";

const getActionsFx = createEffect(() => {
  return getActions();
});

const getActionFx = createEffect((id: number) => {
  return getAction(id);
});

export const $actions = createStore<Action[]>([]).on(getActionsFx.doneData, (_, actionsResponse) => actionsResponse.data);

export const $currentAction = createStore<Action>({
  id: 0,
  name: "",
  pattern: "",
  method: "GET",
  conditions: [],
  schema_ids: [],
  disabled: false,
}).on(getActionFx.doneData, (_, actionResponse) => actionResponse.data);

export const effects = {
  getActionsFx,
  getActionFx,
};

export const events = {};
