import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import {HashRouter} from "react-router";
import Routes from "./routes.jsx";

createRoot(document.getElementById('root')).render(
  <StrictMode>
      <HashRouter>
          <Routes/>
      </HashRouter>
  </StrictMode>,
)
