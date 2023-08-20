import { h, text } from 'forest';
import { withRoute } from 'atomic-router-forest';

import * as routes from '@/shared/routes'
import { Link } from '@/shared/lib/router';

export const HomePage = () => {
  h('div', {
    classList: ['flex', 'flex-col'],
    fn() {
      // This allows to show/hide route if page is matched
      // It is required to call `withRoute` inside `h` call
      withRoute(routes.home);

      text`Hello from the home page`;

      Link(routes.logsList, {
        text: `Show logs list`,
      });
    },
  });
}