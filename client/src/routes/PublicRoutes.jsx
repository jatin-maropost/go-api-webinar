import React, { Suspense } from "react";
import { Switch, Route } from "react-router-dom";

const Home = React.lazy(() => import("../pages/Home"));

const PublicRoutes = (props) => {
  return <Route path="/" exact component={Home} />;
};

export default PublicRoutes;
