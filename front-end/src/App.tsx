import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "./pages/Home";
import Aio from "./pages/Aio";
import Test from "./pages/Test";
import Layout from "./components/Layout";

function App() {
    return (
        <BrowserRouter>
            <Layout>
                <Routes>
                    <Route path="/" element={<Home />} />
                    <Route path="/aio" element={<Aio />} />
                    <Route path="/test" element={<Test />} />
                </Routes>
            </Layout>
        </BrowserRouter>
    );
}

export default App;
