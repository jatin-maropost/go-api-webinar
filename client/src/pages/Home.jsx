import React from "react";
import { Link } from "react-router-dom";

const Home = (props) => {
  return (
    <div>
      <h1>Home here</h1>
      <h2>
        <Link to="/login">Login</Link>
      </h2>
    </div>
  );
};

export default Home;
