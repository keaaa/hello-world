import http from 'k6/http';
import { sleep } from 'k6';

// k6 run --vus 10 --duration 600s loadtest.js
export default function() {
  http.get('https://client-hello-world-dev.playground.radix.equinor.com');
  sleep(1);
}