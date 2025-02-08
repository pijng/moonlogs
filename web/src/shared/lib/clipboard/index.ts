import { createEffect, createEvent, createStore, sample } from "effector";

const CLIPBOARD_MODE_KEY = "clipboard_mode";

type ClipboardMode = "copy" | "no-copy";

export const $clipboardMode = createStore<ClipboardMode>("no-copy");
export const $shouldCopyToClipboard = $clipboardMode.map((mode) => mode === "copy");
export const clipboardModeChanged = createEvent<ClipboardMode>();

export const loadClipboardFromStorageFx = createEffect(() => {
  const clipboardMode: ClipboardMode = localStorage.getItem(CLIPBOARD_MODE_KEY) as ClipboardMode;

  return clipboardMode;
});

const setClipboardModeToStorageFx = createEffect((clipboardMode: ClipboardMode) => {
  if (!clipboardMode) return;

  return localStorage.setItem(CLIPBOARD_MODE_KEY, clipboardMode);
});

sample({
  clock: [clipboardModeChanged, loadClipboardFromStorageFx.doneData],
  target: [setClipboardModeToStorageFx, $clipboardMode],
});
