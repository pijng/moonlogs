import { using } from "forest";
import { onAppMount } from "atomic-router-forest";
import "@/routing";

import { Pages } from "@/pages";
import { appMounted } from "@/routing";

function Application() {
  Pages();
  onAppMount(appMounted);
}

using(document.querySelector("body")!, Application);
