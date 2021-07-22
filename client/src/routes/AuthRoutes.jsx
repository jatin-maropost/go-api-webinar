import React, { Suspense } from "react";
import { Route, Switch } from "react-router-dom";

const Login = React.lazy(() => "../pages/Auth/Login.jsx");
const SignIn = React.lazy(() => "../pages/Auth/SignIn.jsx");

const AuthRoutes = () => {
  return (
    <Suspense fallback={<div>Loading...</div>}>
      <Switch>
        <Route path="/login" component={Login} />
        <Route path="/signin" component={SignIn} />
      </Switch>
    </Suspense>
  );
};

export default AuthRoutes;
