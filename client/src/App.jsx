import React from "react";
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import { PublicRoutes } from "./routes";

const App = (props) => {
  return (
    <Router>
      <PublicRoutes />
    </Router>
  );
};

export default App;
