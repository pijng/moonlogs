import { RouteInstance } from "atomic-router";
import { Effect, sample, Store } from "effector";
import { AnyFormValues, Form } from "effector-forms";

export const manageSubmit = <
  F extends AnyFormValues,
  FX extends F & { id?: number },
  R extends { success: boolean; data: { id?: unknown }; error: string },
>({
  form,
  actionFx,
  error,
  route,
  currentModel,
}: {
  form: Form<F>;
  actionFx: Effect<FX, R, Error>;
  error: Store<string>;
  route: RouteInstance<any>;
  currentModel?: Store<FX>;
}) => {
  if (!!currentModel) {
    sample({
      source: currentModel,
      clock: form.formValidated,
      fn: (currentEntity, entityToEnrich) => {
        return { ...entityToEnrich, id: currentEntity.id } as FX;
      },
      target: actionFx,
    });
  } else {
    sample({
      source: form.formValidated,
      fn: (formData) => ({ ...formData, id: 0 }) as FX,
      target: actionFx,
    });
  }

  sample({
    source: actionFx.doneData,
    filter: (response) => response.success && Boolean(response.data.id),
    target: [form.reset, error.reinit, route.open],
  });

  sample({
    source: actionFx.doneData,
    filter: (response) => !response.success,
    fn: (response) => response.error,
    target: error,
  });
};
