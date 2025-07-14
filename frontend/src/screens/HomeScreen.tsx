import { useState } from "react";
import LinkModel from "../models/link.model";
import { FaCheck, FaCopy } from "react-icons/fa6";
import ReCaptcha from "react-google-recaptcha";
import { RecaptchaSiteKey } from "../config/recaptcha";

export default function HomeScreen() {

    const [error, setError] = useState("");
    const [originalUrl, setOriginalUrl] = useState("");
    const [shortUrl, setShortUrl] = useState("");
    const [isUrlCopied, setUrlCopied] = useState(false);
    const [isCaptchaVerified, setCaptchaVerified] = useState(false);

    async function submit(e: React.FormEvent) {
        e.preventDefault();

        setError("");
        setShortUrl("");

        try {
            if(!isCaptchaVerified) throw "Please complete the captcha verification";

            const url = await LinkModel.create(originalUrl);
            setShortUrl(url);
        }
        catch(ex: any) {
            setError(ex.toString());
        }
    }

    function copyShortUrl(e: React.MouseEvent) {
        e.preventDefault();
        navigator.clipboard.writeText(shortUrl);
        setUrlCopied(true);

        setTimeout(() => {
            setUrlCopied(false);
        }, 2000);
    }

    function onRecaptchaChanged(value: string|null) {
        if(value == null || value == "") {
            setCaptchaVerified(false);
        }
        else {
            setCaptchaVerified(true);
        }
    }

    return(
        <div className="min-h-dvh w-dvw bg-zinc-900 flex justify-center items-center">
            
            <div className="w-[95%] max-w-[600px] flex flex-col p-10 gap-10">
                <div className="flex flex-col justify-center items-center">
                    <h1 className="text-5xl font-bold text-white">ShortenIt</h1>
                    <p className="text-xl text-zinc-200">Easy to use URL Shortener</p>
                </div>

                <div>
                    { (error != "") ? <p className="text-red-500 mb-3 text-lg">{error}</p> : <></> }

                    <form onSubmit={submit} className="flex gap-3">
                        <input value={originalUrl} onChange={e => {
                            setOriginalUrl(e.target.value);
                        }} className="text-zinc-100 w-full border px-3 h-[45px] text-lg rounded-lg border-zinc-600 flex-1" type="text" placeholder="Enter URL here..." required />

                        <button className="text-lg text-white bg-teal-500 h-[45px] px-5 rounded-lg cursor-pointer">Shorten It!</button>
                    </form>
                </div>
                
                {
                    (shortUrl != "") ? <div onClick={copyShortUrl} className="border border-teal-500 bg-teal-500/10 p-5 rounded-lg cursor-pointer flex items-center justify-between">
                        <div>
                            <h4 className="text-zinc-300">Here's your shortened url:</h4>
                            <p className="text-white text-xl">{ shortUrl }</p>
                        </div>

                        {
                            (!isUrlCopied) ? <FaCopy size={20} className="text-white mr-1" /> : <FaCheck size={20} className="text-white mr-1" /> 
                        }
                    </div> : <></>
                }

                <div className="flex justify-center items-center">
                    <ReCaptcha sitekey={RecaptchaSiteKey} onChange={onRecaptchaChanged} />
                </div>

            </div>

        </div>
    );
}