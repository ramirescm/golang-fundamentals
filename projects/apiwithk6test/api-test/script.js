import http from 'k6/http';
import { check, sleep } from 'k6';

export let options = {
    stages: [
        { duration: '30s', target: 10 },  // Ramp-up to 10 virtual users over 30 seconds
        { duration: '1m', target: 10 },   // Stay at 10 virtual users for 1 minute
        { duration: '30s', target: 0 },   // Ramp-down to 0 virtual users over 30 seconds
    ],
};

export default function () {
    let res = http.get('http://localhost:8080/hello');
    check(res, {
        'status is 200': (r) => r.status === 200,
    });

    // Simulate a user "thinking" by sleeping for a short duration
    sleep(5);
}