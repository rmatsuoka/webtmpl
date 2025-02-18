async function getCount() {
  const res = await fetch("/api/count");
  const body = await res.json();
  if (!res.ok) {
    throw new Error(body["message"]);
  }
  return body["count"];
}

async function countup() {
  const res = await fetch("/api/count", {
    method: "POST",
  });
  const body = await res.json();
  if (!res.ok) {
    throw new Error(body["message"]);
  }
  return body["count"];
}

const count = document.querySelector("#count");

function setCount(c) {
  count.textContent = c;
}

function main() {
  getCount().then((c) => setCount(c));

  const form = document.querySelector("form#counter");
  form.addEventListener("submit", (e) => {
    e.preventDefault();
    countup().then((c) => setCount(c));
  });
}

main();
