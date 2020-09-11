import React, {Fragment, Component} from 'react';
import MainInput from "./mainFormElements/mainInput";
import {signUp} from "../../Js Functionals/MainPage/IOMethods/signUpMethods";
import '../../CSS Designs/extra-css.css'

class MainSignUpMenu extends Component {
    constructor(props, context) {
        super(props, context);
    }

    render() {
        return (
            <div id={this.props.id} style={this.props.style} className="content Login-SignUp-Menu">
                <div className="ui form formPadding ui-rtl">

                    <MainInput id='SignUp-UserName' maxLength='30' textName='نام کاربری' placeHolder='Username'/>

                    <MainInput id="SignUp-FirstName" maxLength='30' textName='نام' placeHolder='FirstName' />

                    <MainInput id="SignUp-LastName" maxLength='30' textName='نام خانوادگی' placeHolder='LastName' />

                    <MainInput id="SignUp-Email" maxLength='50' textName='ایمیل' placeHolder='Email address'/>

                    <MainInput id="SignUp-Password" maxLength='30' textName='رمز عبور' placeHolder='Password'
                               isPassword={true}/>

                    <MainInput id="SignUp-RepeatPassword" maxLength='30' textName='تکرار رمز عبور'
                               placeHolder='Repeat Password' isPassword={true}/>

                    <div className="ui form formPadding">
                        <div className="ui paragraph paragraphInput">نوع</div>
                        <label>
                            <select className="ui dropdown" id="signUpKind">
                                <option value="employer">کارفرما</option>
                                <option value="freelancer">فریلنسر</option>
                            </select>
                        </label>
                    </div>

                    <div className="ui form formPadding">
                        <label>
                            <input onClick={() => {
                                signUp(this.props.onClose)
                            }} type="submit" id="signUpButton" className="ui green button"
                                   value="ثبت نام"/>
                        </label>
                    </div>
                </div>
            </div>
        );
    }
}

export default MainSignUpMenu;
