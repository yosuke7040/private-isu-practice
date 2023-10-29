import { SharedArray } from 'k6/data';

const accounts = new SharedArray("accounts", function() {
  return JSON.parse(open('./accounts.json'));
});

// export default function() {
//   return accounts[Math.floor(Math.random() * accounts.length)];
// };
export function getAccount() {
  return accounts[Math.floor(Math.random() * accounts.length)];
}