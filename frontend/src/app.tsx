import { MetaProvider, Title } from "@solidjs/meta";
import { Router } from "@solidjs/router";
import { FileRoutes } from "@solidjs/start/router";
import { Suspense, createResource } from "solid-js";
import "./app.css";

const fetchMessage = async () => {
  const res = await fetch("http://127.0.0.1:8000/api/hello/");
  return res.json();
};

export default function App() {
  const [data] = createResource(fetchMessage);
  return (
    <Router
      root={props => (
        <MetaProvider>
          <Title>SolidStart - Basic</Title>
          <a href="/">Index</a>
          <a href="/about">About</a>
          <p>{data()?.message}</p>
          <Suspense>{props.children}</Suspense>
        </MetaProvider>
      )}
    >
      <FileRoutes />
    </Router>
  );
}
