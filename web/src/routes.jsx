import React from "react";
import { Route, Routes as R } from "react-router";
import HomePage from "./pages/HomePage.jsx";
import NotFoundPage from "./pages/NotFoundPage.jsx";
import PanelLayout from "./components/PanelLayout.jsx";
import ServiceEdit from "./pages/services/ServiceEdit.jsx";
import ServiceNew from "./pages/services/ServiceNew.jsx";
import ServiceList from "./pages/services/ServiceList.jsx";
import RouteList from "./pages/routes/RouteList.jsx";
import RouteNew from "./pages/routes/RouteNew.jsx";
import RouteEdit from "./pages/routes/RouteEdit.jsx";

function Routes() {
  return (
    <R>
      <Route element={<PanelLayout />}>
        <Route index element={<HomePage />} />
        <Route path="config">
          <Route path="services">
            <Route index element={<ServiceList />} />
            <Route path="new" element={<ServiceNew />} />
            <Route path=":id/edit" element={<ServiceEdit />} />
          </Route>
          <Route path="routers">
            <Route index element={<RouteList />} />
            <Route path="new" element={<RouteNew />} />
            <Route path=":id/edit" element={<RouteEdit />} />
          </Route>
        </Route>
      </Route>

      <Route path="*" element={<NotFoundPage />} />
    </R>
  );
}

export default Routes;
