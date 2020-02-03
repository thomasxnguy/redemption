import React, { Component } from "react";
import { HashRouter, Route, Switch } from "react-router-dom";
import { ToastProvider } from "react-toast-notifications";
// import { renderRoutes } from 'react-router-config';
import "./App.scss";

const loading = () => (
  <div className="animated fadeIn pt-3 text-center">Loading...</div>
);

// Containers
const DefaultLayout = React.lazy(() => import("./containers/DefaultLayout"));

// Pages
const Page404 = React.lazy(() => import("./Page404"));

class App extends Component {
  render() {
    return (
      <ToastProvider placement="bottom-right" autoDismissTimeout="3000">
        <HashRouter>
          <React.Suspense fallback={loading()}>
            <Switch>
              <Route
                exact
                path="/404"
                name="Page 404"
                render={props => <Page404 {...props} />}
              />
              <Route
                path="/"
                name="Home"
                render={props => <DefaultLayout {...props} />}
              />
            </Switch>
          </React.Suspense>
        </HashRouter>
      </ToastProvider>
    );
  }
}

export default App;
