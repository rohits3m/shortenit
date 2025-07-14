import { BrowserRouter, Route, Routes } from "react-router-dom";
import HomeScreen from "./screens/HomeScreen";

export default function App() {
    return(
        <BrowserRouter>
            <Routes>

                <Route path="/" element={ <HomeScreen /> } />

            </Routes>
        </BrowserRouter>
    );
}