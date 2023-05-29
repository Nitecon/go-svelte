export function updateInput(inputName) {
    const inputElement = document.querySelector(`[name="${inputName}"]`);
    if (inputElement) {
        inputElement.dispatchEvent(new Event('input'));
    }
}
function uc(str) {
    return str.charAt(0).toUpperCase() + str.slice(1);
}
export function matchField(inputName) {
    let niceName = uc(inputName);
    return function (value) {
        if (value === undefined || value === null || value === '') {
            return null;
        }
        return value === document.querySelector(`[name="${inputName}"]`).value ? null : {matchField: `${niceName} does not match!`};
    };
}
export function email() {
    return function (value) {
        if (value === undefined || value === null) {
            return {validateEmail: "Email cannot be null or undefined!"};
        }

        // Simple regex pattern for email validation
        let pattern = /\S+@\S+\.\S+/;

        return pattern.test(value) ? null : {validateEmail: "Must be a valid email address!"};
    };
}
export function minLength(amount) {
    return function (value) {
        if (value === undefined || value === null) {
            return {minLength: "Input cannot be null or undefined!"};
        }
        // The length of an empty string is 0, so it will be correctly handled
        return value.length >= amount ? null : {minLength: "Must be at least " + amount + " characters long!"};
    };
}
export function minUpperCase(amount) {
    return function (value) {
        if (value === undefined || value === null || value === '') {
            return null;
        }
        let upperCase = value.match(/[A-Z]/g);
        return upperCase && upperCase.length >= amount ? null : {minUpperCase: "Must contain at least ( " + amount + " ) uppercase characters!"};
    };
}

export function minLowerCase(amount) {
    return function (value) {
        if (value === undefined || value === null || value === '') {
            return null;
        }
        let lowerCase = value.match(/[a-z]/g);
        return lowerCase && lowerCase.length >= amount ? null : {minLowerCase: "Must contain at least ( " + amount + " ) uppercase characters!"};
    };
}

export function minNumbers(amount) {
    return function (value) {
        if (value === undefined || value === null || value === '') {
            return null;
        }
        let numbers = value.match(/[0-9]/g);
        return numbers && numbers.length >= amount ? null : {minNumbers: "Must contain at least ( " + amount + " ) numbers!"};
    };
}

export function minSymbols(amount) {
    return function (value) {
        if (value === undefined || value === null || value === '') {
            return null;
        }
        let symbols = value.match(/[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]/g);
        return symbols && symbols.length >= amount ? null : {minSymbols: "Must contain at least ( " + amount + " ) symbols!"};
    };
}

export default function validatePassword(minUpperCase, minLowerCase, minSymbols,minNumbers, length=8 ) {
    return function(value){
        if (value === undefined || value === null || value === '') {
            return {validatePassword: "Password is required!"};
        }
        let upperCase = value.match(/[A-Z]/g);
        let lowerCase = value.match(/[a-z]/g);
        let numbers = value.match(/[0-9]/g);
        let symbols = value.match(/[!@#$%^&*()_+\-=\[\]{};':"\\|,.<>\/?]/g);

        if ((!upperCase || upperCase.length < minUpperCase) ||
            (!lowerCase || lowerCase.length < minLowerCase) ||
            (!numbers || numbers.length < minNumbers) ||
            (!symbols || symbols.length < minSymbols)) {
            let errorMessage = [];
            if (!upperCase || upperCase.length < minUpperCase) {
                errorMessage.push("minimum (" + minUpperCase + ") uppercase");
            }
            if (!lowerCase || lowerCase.length < minLowerCase) {
                errorMessage.push("minimum (" + minLowerCase + ") lowercase");
            }
            if (!numbers || numbers.length < minNumbers) {
                errorMessage.push("minimum (" + minNumbers + ") numbers");
            }
            if (!symbols || symbols.length < minSymbols) {
                errorMessage.push("minimum (" + minSymbols + ") symbols");
            }
            if (value.length < length) {
                errorMessage.push("minimum (" + length + ") characters total");
            }
            return {validatePassword: errorMessage.join(', ')};
        }
        return null;
    }
}

