import React from "react";

const Dashboard = React.lazy(() => import("./views/Dashboard"));
const AllCodes = React.lazy(() => import("./pages/codes/AllCodes"));
const CreateCodes = React.lazy(() => import("./pages/codes/CreateCodes"));
const ViewCode = React.lazy(() => import("./pages/codes/ViewCode"));

// https://github.com/ReactTraining/react-router/tree/master/packages/react-router-config
const routes = [
  { path: "/", exact: true, name: "Home" },
  { path: "/dashboard", name: "Dashboard", component: Dashboard },
  { path: "/codes", exact: true, name: "Codes", component: AllCodes },
  { path: "/codes/all", name: "All", component: AllCodes },
  { path: "/codes/create", name: "New", component: CreateCodes },
  { path: "/codes/:id", name: "View", component: ViewCode }
];

export default routes;
