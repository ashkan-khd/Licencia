import {httpExcGET, httpGet} from "../AlphaAPI";
import {mainPageName} from "../FileNames";
import Cookies from 'js-cookie';
import {
    changePasswordUrlEmployer,
    changePasswordUrlFreeLancer,
    gitHubUrl, saveGithubUrlFreeLancer,
    saveProfileUrlEmployer,
    saveProfileUrlFreeLancer,
    urlGetEmployerProfileInfo,
    urlGetFreelancerProfileInfo
} from "../urlNames";


let usernameField
let shownNameField
let firstNameField
let lastNameField
let emailField
let siteAddressField
let telephoneNumberField
let gitHubAccountField
let descriptionField
let addressField

let username;
let password;
let shownName;
let firstname;
let lastname;
let email;
let siteAddress;
let telephoneNumber;
let gitHubAccount;
let gitHubRepo;
let description;
let address;
let projectsId;
let requestedProjectsId;


function fillForProfileFields() {
    usernameField = document.getElementById("usernameField");
    shownNameField = document.getElementById('showingNameField')
    firstNameField = document.getElementById("firstNameField");
    lastNameField = document.getElementById("lastNameField");
    emailField = document.getElementById("emailField");
    siteAddressField = document.getElementById("siteAddressField");
    telephoneNumberField = document.getElementById("telephoneNumberField");
    gitHubAccountField = document.getElementById("githubAccountField");
    descriptionField = document.getElementById("descriptionField");
    addressField = document.getElementById("addressField");
}

function logOut() {
    // Todo
    window.location.href = mainPageName;
}

const gitHubAccountPart = document.getElementById("gitHubAccountPart");

function initGithubRepos() {
    fill();
    firstRepoDiv.style.display = 'none';
    secondRepoDiv.style.display = 'none';
    thirdRepoDiv.style.display = 'none';
    iconDiv.style.display = 'none';
}

export let isFreeLancer = true;

export function initTransitionsStart(profileTransition, ...transitions) {
    alert('transitions: ' + transitions)
    // transitions.forEach((value => value.toggleVisibility()))
}

export function loadProfileMenu() {
    fillForProfileFields()
    // alert('IsFreeLancer: ' + Cookies.get('isfreelancer'))
    isFreeLancer = Cookies.get('isfreelancer');
    alert('Cookies: "' + isFreeLancer + "'")
    if (!isFreeLancer) {
        httpGet(urlGetEmployerProfileInfo, {
            'Content-Type': 'application/json',
            'Token': Cookies.get('auth')
        }, handleSuccessGetProfileInfo, handleDenyGetProfileInfo);
    } else {
        httpGet(urlGetFreelancerProfileInfo, {
            'Content-Type': 'application/json',
            'Token': Cookies.get('auth')
        }, handleSuccessGetProfileInfo, handleDenyGetProfileInfo);
    }
    initGithubRepos();
}

function handleSuccessGetProfileInfo(value) {
    let messages = value;
    alert(JSON.stringify(messages));
    username = messages.username;
    shownName = messages['shown-name']
    firstname = messages['first-name'];
    lastname = messages['last-name'];
    email = messages.email;
    description = messages.description;
    alert('description: ' + messages.description)
    telephoneNumber = messages['phone-number'];
    address = messages.address;
    projectsId = messages['project-ids'];
    fillCommonFields();
    if (isFreeLancer) {
        gitHubAccount = messages.github;
        gitHubRepo = messages['github-repos'];
        siteAddress = messages.website;
        requestedProjectsId = messages['req-project-ids'];
        fillFreelancerSpecialFields();
    }
}

function fillFreelancerSpecialFields() {
    siteAddressField.value = siteAddress;
    //TODO : remaining;
}

function fillCommonFields() {
    usernameField.value = username;
    shownNameField.value = shownName;
    firstNameField.value = firstname;
    lastNameField.value = lastname;
    emailField.value = email;
    telephoneNumberField.value = telephoneNumber;
    addressField.value = address;
    descriptionField.value = description;
}

function handleDenyGetProfileInfo(value) {
    alert(JSON.stringify(value));
    console.log("raft too handleDenyGetProfileInfo");
}

export let repoDiv
export let iconDiv
export let repoInput
export let firstRepoDiv
export let secondRepoDiv
export let thirdRepoDiv
export let firstRepoLink
export let secondRepoLink
export let thirdRepoLink
export let githubAccountField
export let gitHubReposDiv

export function fill() {
    repoDiv = document.getElementById("addRepoDiv");
    iconDiv = document.getElementById("plusRepoIconDiv");
    repoInput = document.getElementById("addRepoInput");
    firstRepoDiv = document.getElementById("firstRepo");
    secondRepoDiv = document.getElementById("secondRepo");
    thirdRepoDiv = document.getElementById("thirdRepo");
    firstRepoLink = document.getElementById("linkRepo1");
    secondRepoLink = document.getElementById("linkRepo2");
    thirdRepoLink = document.getElementById("linkRepo3");
    githubAccountField = document.getElementById("githubAccountField");
    gitHubReposDiv = document.getElementById("gitHubRepos");

}

export function openAddRepoDiv() {
    fill()
    let counter = 0;
    if (firstRepoDiv.style.display !== "none") counter += 1;
    if (secondRepoDiv.style.display !== "none") counter += 1;
    if (thirdRepoDiv.style.display !== "none") counter += 1;
    if (counter === 3) {
        alert("you can have only 3 repository")
        return;
    }
    repoDiv.style.display = "block";
    iconDiv.style.display = "none";
    repoInput.value = "";
    repoInput.focus();
}


export function closeAddRepoDiv() {
    fill()
    if (repoInput.value === "") {
        // $ep
        repoDiv.style.display = "none";
        iconDiv.style.display = "block";
        return;
    }
    if (firstRepoDiv.style.display === "none") {
        showRepo(firstRepoDiv, firstRepoLink, repoInput.value)
    } else if (secondRepoDiv.style.display === "none") {
        showRepo(secondRepoDiv, secondRepoLink, repoInput.value)
    } else if (thirdRepoDiv.style.display === "none") {
        showRepo(thirdRepoDiv, thirdRepoLink, repoInput.value)
    } else {
        alert("you can have only 3 repository")
        return;
    }
    repoDiv.style.display = "none";
    iconDiv.style.display = "block";
    repoInput.value = "";
    let counter = 0;
    if (firstRepoDiv.style.display !== "none") counter += 1;
    if (secondRepoDiv.style.display !== "none") counter += 1;
    if (thirdRepoDiv.style.display !== "none") counter += 1;
    if (counter === 3) {
        iconDiv.style.display = "none";
    }
}

function showRepo(repoDiv, repoLink, text) {
    let textNode = document.createTextNode(text);
    removeAllChild(repoLink);
    repoLink.appendChild(textNode);
    repoLink.href = gitHubUrl + gitHubAccountField.value + '/' + text
    repoDiv.style.display = "block";
}

export function removeRepo(element) {
    fill()
    element.style.display = "none";
    iconDiv.style.display = "block";
}

function removeAllChild(element) {
    while (element.firstChild) {
        element.removeChild(element.firstChild);
    }
}

export function accountGithubChanged() {
    fill()
    if (githubAccountField.value === "") {
        firstRepoDiv.style.display = "none";
        secondRepoDiv.style.display = "none";
        thirdRepoDiv.style.display = "none";
        repoDiv.style.display = "none";
        iconDiv.style.display = "none";
    } else {
        iconDiv.style.display = "block";
    }
}

const MainProfileTransition = 'fade up';
export const mainProfileContent = document.getElementById('MainProfileContent');
export const profile = document.getElementById('profile');
export const gitHubRepoContent = document.getElementById('githubReposContent');
export const changePasswordContent = document.getElementById('changingPasswordContent');

export function changeMainProfileContent(content) {
    /*let showingDisplay = getShowingDisplay();
    if (showingDisplay != null && content.id !== showingDisplay.id) {
        $('#' + showingDisplay.id).transition(MainProfileTransition);
        $('#' + content.id).transition(MainProfileTransition);
    }*/
}

function getShowingDisplay() {
    for (let childElement of mainProfileContent.children) {
        console.log('showingDisplay: ' + childElement.id)
        console.log('****: ' + childElement.style.display)
        if (childElement.style.display != '' && childElement.style.display !== 'none') {
            return childElement;
        }
    }
    return null;
}


function modal(modalId, command) {
    /*if (document.getElementById(modalId) != null) {
        $('#' + modalId)
            .modal(command);
    }*/
}

function successSaveProfile(value) {
    alert('Profile Saved Successfully')
    // eslint-disable-next-line no-restricted-globals
    location.reload();
}

function errorSaveProfile(value) {
    //Error Handling
    alert('We Have An Error')
    alert('error: ' + value.message)
}

export function saveProfile() {
    let getValue = (firstValue, secondValue) => secondValue == null ? firstValue : secondValue;
    const data = {
        'shown-name': getValue(shownName, shownNameField.value),
        'first-name': getValue(firstname, firstNameField.value),
        'last-name': getValue(lastname, lastNameField.value),
        'phone-number': telephoneNumberField.value,
        'address': addressField.value,
        'description': descriptionField.value
    }
    alert('data: ' + JSON.stringify(data))
    telephoneNumber = telephoneNumberField.value;
    address = addressField.value;
    description = descriptionField.value;
    httpExcGET('post', isFreeLancer ? saveProfileUrlFreeLancer : saveProfileUrlEmployer,
        data, successSaveProfile, errorSaveProfile, {
            'Token': Cookies.get('auth')
        })
}

export function submitGitPart() {
    fill()
    if (isFreeLancer) {
        let gitLinks = [];
        let size = 0;
        if (firstRepoDiv.style.display !== "none") {
            gitLinks[size] = document.getElementById('linkRepo1').text();
            size += 1;
        }
        if (secondRepoDiv.style.display !== "none") {
            gitLinks[size] = document.getElementById('linkRepo2').text();
            size += 1;
        }
        if (thirdRepoDiv.style.display !== "none") {
            gitLinks[size] = document.getElementById('linkRepo3').text();
            size += 1;
        }
        let data = {
            'website': siteAddressField.value,
            'github-repos': gitLinks,
            'github': githubAccountField.value
        }
        siteAddress = siteAddressField.value;
        gitHubAccount = githubAccountField.value;
        let headers = {
            'Content-Type': 'application/json',
            'token': Cookies.get('auth')
        }
        httpExcGET('POST', saveGithubUrlFreeLancer, data, successGithubPartSubmit, denyGithubPartSubmit, headers);
    }
}

function successGithubPartSubmit(value) {
    alert("post successfully" + " value : " + JSON.stringify(value));
}

function denyGithubPartSubmit(value) {
    alert("post deny" + " value : " + JSON.stringify(value));
}

const oldPasswordField = document.getElementById("oldPasswordField");
const newPasswordField = document.getElementById("passwordField");
const repeatNewPasswordField = document.getElementById("repeatPasswordField");
export function changePassword() {
    if (oldPasswordField.value === "" || newPasswordField.value === "" || repeatNewPasswordField.value === "") {
        alert("you have empty field")
    } else {
        if (newPasswordField.value !== repeatNewPasswordField.value) {
            alert("passwords doesn't match")
        } else {
            if (oldPasswordField.value !== password) {
                alert("old password is incorrect")
            } else {
                let data = {
                    'new-password': siteAddressField.value,
                }
                let headers = {
                    'Content-Type': 'application/json',
                    'token': Cookies.get('auth')
                }
                password = newPasswordField.value;
                httpExcGET('POST', isFreeLancer ? changePasswordUrlFreeLancer : changePasswordUrlEmployer,
                    data, successChangePassword, denyChangePassword, headers)
            }
        }
    }
}

function successChangePassword(value) {
    alert("password changed successfully" + "  value : " + JSON.stringify(value))
}

function denyChangePassword(value) {
    alert("password doesn't change" + "  value : " + JSON.stringify(value))
}


function openClose() {
    /*$('.ui.sidebar')
        .sidebar('toggle')
    ;*/
}


let profileComponent;
let linksComponent;

export function switchProfileToLinks() {
    fillComponents();
    linksComponent.style.display = "block"
    profileComponent.style.display = "none";
}

export function switchLinksToProfile() {
    fillComponents();
    linksComponent.style.display = "none"
    profileComponent.style.display = "block";
}

function fillComponents() {
    profileComponent = document.getElementById('profileComponent')
    linksComponent = document.getElementById('linksComponent')
}
