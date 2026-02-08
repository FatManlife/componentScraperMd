import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "./pages/Home";
import Aio from "./pages/Aio";
import Test from "./pages/Test";
import Layout from "./components/Layout";
import Case from "./pages/Case";
import Cooler from "./pages/Cooler";
import Cpu from "./pages/Cpu";
import Fan from "./pages/Fan";
import Gpu from "./pages/Gpu";
import Hdd from "./pages/Hdd";
import Laptop from "./pages/Laptop";
import Motherboard from "./pages/Motherboard";
import Pc from "./pages/Pc";
import PcMini from "./pages/PcMini";
import Psu from "./pages/Psu";
import Ram from "./pages/Ram";
import Ssd from "./pages/Ssd";

function App() {
    return (
        <BrowserRouter>
            <Layout>
                <Routes>
                    <Route path="/" element={<Home />} />
                    <Route path="/aio" element={<Aio />} />
                    <Route path="/test" element={<Test />} />
                    <Route path="/case" element={<Case />} />
                    <Route path="/cooler" element={<Cooler />} />
                    <Route path="/cpu" element={<Cpu />} />
                    <Route path="/fan" element={<Fan />} />
                    <Route path="/gpu" element={<Gpu />} />
                    <Route path="/hdd" element={<Hdd />} />
                    <Route path="/laptop" element={<Laptop />} />
                    <Route path="/motherboard" element={<Motherboard />} />
                    <Route path="/pc" element={<Pc />} />
                    <Route path="/pcmini" element={<PcMini />} />
                    <Route path="/psu" element={<Psu />} />
                    <Route path="/ram" element={<Ram />} />
                    <Route path="/ssd" element={<Ssd />} />
                </Routes>
            </Layout>
        </BrowserRouter>
    );
}

export default App;
