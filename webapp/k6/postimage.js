import http from "k6/http";
import { check } from "k6";
import { url } from "./config.js";
import { getAccount } from "./accounts.js";
import { parseHTML } from "k6/html";

const testImage = open("testimage.jpg", "b");

export default function() {
  const account = getAccount();
  const res = http.post(url("/login"), {
    account_name: account.name,
    password: account.password,
  });
  const doc = parseHTML(res.body);
  const token = doc.find('input[name="csrf_token"]').attr('value');

  http.post(url("/"), {
    file: http.file(testImage, "testimage.jpg", "image/jpeg"),
    csrf_token: token,
    body: "posted by k6",
  });
}
