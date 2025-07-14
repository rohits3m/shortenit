import { Axios } from "axios";

export interface IApiResponse {
    success: boolean;
    data: any;
    message?: string;
}

const Api = new Axios({
    baseURL: "http://localhost:5000/v1",
    headers: {
        "Content-Type": "application/json"
    }
});

export default Api;