import http from "k6/http";

const url = `${__ENV.TARGET_URL}/`;
const payload = JSON.stringify({});
const params = {};

export default function () {
    http.post(url, payload, params);
}
