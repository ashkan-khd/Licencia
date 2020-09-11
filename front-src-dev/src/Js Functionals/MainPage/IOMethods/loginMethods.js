import MainTextField from "../../../Components/MainPageComponents/mainTextField";
import {httpExcGET} from "../../AlphaAPI";
import {urlLogin} from "../../urlNames";
import Cookies from "js-cookie";
import {goToPage} from "../../PageRouter";
import {profilePagePath} from "../../PagePaths";
import {loginInvalidPasswordLabel, loginNotSignedUpKeypointLabel} from "../ioErrors";
import {emptyFieldsFromErrors, hasEmpty} from "./Utils";

let loginKeypoint;
let loginPassword;
let loginKind;
let loginCloseModalFunc;
function setLoginFields() {
    loginKeypoint = document.getElementById("login-KeyPoint");
    loginPassword = document.getElementById("login-Password");
    loginKind = document.getElementById("loginKind");
}

export function emptyLoginFields() {
    setLoginFields()
    loginKeypoint.value = "";
    loginPassword.value = "";
    loginKind.value = "";
}
export let emptyLoginFieldsFromErrors = () => emptyFieldsFromErrors(loginKeypoint, loginPassword)

export function login(func) {
    loginCloseModalFunc = func;
    setLoginFields()
    let doc = hasEmpty(loginKeypoint, loginPassword);
    emptyFieldsFromErrors(loginKeypoint, loginPassword)
    if (doc != null) {
        MainTextField.setFieldError(doc, true)
        MainTextField.showErrorLabel(doc, 'fill It, Dude')
        // setTimeout(() => alert("fill the red box!!"), 1000);
    } else {
        const data = {
            id: loginKeypoint.value,
            password: loginPassword.value
        }
        alert('data: ' + JSON.stringify(data))
        const promise = httpExcGET('post', urlLogin, data, handleSuccessLogin, handleErrorLogin, {
            'Content-Type': 'application/json'
        }, {
            key: 'account-type',
            value: loginKind.value
        });

    }
}

function handleSuccessLogin(value) {
// todo go to Profile Menu And Save Auth
    alert("Login Successful")
    Cookies.set("isfreelancer", loginKind.value === "freelancer");
    goToPage(profilePagePath);
    emptyLoginFields();
}

function handleErrorLogin(value) {
    // todo error the fields
    alert("Login Failed")
    alert('Server Message: ' + value.message)

    switch (value.message) {
        case 'not signed up username':
        case 'not signed up email':
            MainTextField.setFieldError(loginKeypoint)
            MainTextField.setFieldError(loginPassword)
            MainTextField.showErrorLabel(loginKeypoint, loginNotSignedUpKeypointLabel)
            break;
        case 'invalid password':
            MainTextField.setFieldError(loginPassword)
            MainTextField.showErrorLabel(loginPassword, loginInvalidPasswordLabel)
            break;
        default:
            alert("Haven't Handled That Error Before");
            console.log("messageError: '" + value.message + "'")
    }
}