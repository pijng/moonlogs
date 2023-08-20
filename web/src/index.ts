import { sample, createEvent } from 'effector';
import { using } from 'forest';
import { createBrowserHistory } from 'history';
import { createHistoryRouter } from 'atomic-router';
import { linkRouter, onAppMount } from 'atomic-router-forest';

import { ROUTES, Pages } from '@/pages';
import { Link } from '@/shared/lib/router';

// Create history instance and router instance to control routing in the app
const history = createBrowserHistory();
const router = createHistoryRouter({ routes: ROUTES });

// This event need to setup initial configuration. You can move it into src/shared
const appMounted = createEvent();

// Attach history for the router on the app start
sample({
  clock: appMounted,
  fn: () => history,
  target: router.setHistory,
});

// Add router into the Link instance to easily resolve routes paths
linkRouter({
  clock: appMounted,
  router,
  Link,
});

function Application() {
  Pages();
  onAppMount(appMounted);
}

using(document.querySelector('body')!, Application);