import { Routes, Route, BrowserRouter } from "react-router-dom";
import { ToastContainer } from "react-toastify";
import "react-toastify/dist/ReactToastify.css";
import Login from "./Components/Login";
import Register from "./Components/Register";
import HomePage from "./Pages/HomePage";
import LibraryPage from "./Pages/LibraryPage";
import NotFound from "./Pages/NotFoundPage";
import PreviewPaperPage from "./Pages/PreviewPaperPage";
import WritePaperPage from "./Pages/WritePaperPage";

function App() {
  return (
    <div className="App">
      <ToastContainer
        position="top-right"
        autoClose={5000}
        hideProgressBar={false}
        newestOnTop={false}
        closeOnClick
        rtl={false}
        pauseOnFocusLoss
        draggable
        pauseOnHover
      />
      <BrowserRouter>
        <Routes>
          <Route exact path="/login" element={<Login />} />
          <Route exact path="/register" element={<Register />} />
          <Route exact path="/home" element={<HomePage />} />
          <Route exact path="/write" element={<WritePaperPage />} />
          <Route exact path="/edit/:paperId" element={<WritePaperPage />} />
          <Route exact path="/library" element={<LibraryPage />} />
          <Route exact path="/paper/:paperId" element={<PreviewPaperPage />} />
          <Route path="*" element={<NotFound />} />
        </Routes>
      </BrowserRouter>
    </div>
  );
}

export default App;
