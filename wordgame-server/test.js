import http from "k6/http";

export default function() {
    let response = http.get("http://10.0.1.2:8088/new");
};
