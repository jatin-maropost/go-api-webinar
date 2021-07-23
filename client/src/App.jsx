import React, { Suspense } from "react";
import { PublicRoutes, AuthRoutes } from "./routes";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import "bootstrap/dist/css/bootstrap.min.css";

const App = (props) => {
  return (
    <Router>
      <Switch>
        <Suspense fallback={<div>Loading...</div>}>
          <PublicRoutes />
          <AuthRoutes />
        </Suspense>
      </Switch>
    </Router>
  );
};

export default App;
