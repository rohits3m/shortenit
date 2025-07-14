import Api, { type IApiResponse } from "../config/api";

export default class LinkModel {

    static async create(url: string): Promise<string> {
        try {
            const response = await Api.post("/link/create", JSON.stringify({ "original_url": url }));
            const json = JSON.parse(response.data) as IApiResponse;

            if(!json.success) throw json.message ?? "Unexpected error occured";

            return json.data as string;
        }
        catch(ex) {
            throw ex;
        }
    }

}