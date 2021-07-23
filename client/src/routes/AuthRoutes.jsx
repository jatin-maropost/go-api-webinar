import React, { Fragment } from "react";
import { Route } from "react-router-dom";

const Login = React.lazy(() => import("../pages/Auth/Login.jsx"));
const SignIn = React.lazy(() => import("../pages/Auth/SignIn.jsx"));

const AuthRoutes = () => {
  return (
    <Fragment>
      <Route path="/login" component={Login} />
      <Route path="/signin" component={SignIn} />
    </Fragment>
  );
};

export default AuthRoutes;
