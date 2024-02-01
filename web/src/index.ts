import { h, spec, using } from "forest";
import { onAppMount } from "atomic-router-forest";
import "@/routing";

import { Pages } from "@/pages";
import { appMounted } from "@/routing";
import { userModel } from "./entities/user";

function Application() {
  h("body", () => {
    spec({
      classList: {
        "w-full": true,
        "min-h-screen": true,
        "bg-white": true,
        "text-slate-700": true,
        "dark:bg-slate-900": true,
        "dark:text-slate-300": true,
        dark: userModel.$currentTheme.map((theme) => theme === "dark"),
      },
    });

    Pages();
    onAppMount(appMounted);
  });
}

using(document.querySelector("html")!, Application);
