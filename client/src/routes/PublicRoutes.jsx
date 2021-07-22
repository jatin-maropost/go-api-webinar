import React, { Suspense } from "react";
import { Switch, Route } from "react-router-dom";

const Home = React.lazy(() => import("../pages/Home"));

const PublicRoutes = (props) => {
  return (
    <Suspense fallback={<div>Loading...</div>}>
      <Switch>
        <Route path="/" component={Home} />
      </Switch>
    </Suspense>
  );
};

export default PublicRoutes;
