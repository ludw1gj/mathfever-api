import { isValidBinary, isValidDecimal, isValidHexadecimal } from './validate';

/**
 * Create an object from a HTMLFormElement if form values are valid (binary, decimal, hex).
 * @param {HTMLFormElement} form - The form to serialize.
 * @param {Function} onError - The function to execute if form is invalid.
 * @returns {object | null} - An object if form is valid or null if invalid.
 */
export const createSerializedAndValidatedForm = (form, onError) => {
  const serializedForm = createSerializedForm(form);
  for (const key in serializedForm) {
    // removes all whitespace from the form data
    serializedForm[key] = serializedForm[key].replace(/\s/g, '');

    // check if key's value is valid
    if (!validateAndParseFormKey(key, serializedForm, onError)) {
      return null;
    }
  }
  // form is valid
  return serializedForm;
};

/**
 * Create an object from a HTMLFormElement.
 * @param {HTMLFormElement} form - The form to serialize.
 * @returns {object} - A serialized form.
 */
const createSerializedForm = form => {
  const elements = form.elements;
  const obj = {}; // { [k: string]: any }

  for (let i = 0; i < elements.length; i += 1) {
    const element = elements[i];
    const type = element.type;
    const name = element.name;
    const value = element.value;

    switch (type) {
      case 'hidden':
      case 'text':
        obj[name] = value;
        break;
      default:
        break;
    }
  }
  return obj;
};

/**
 * Check if value of key in serialized form is valid.
 * @param {string} key - The key of the value to validate.
 * @param {HTMLFormElement} serializedForm - A serialize form.
 * @param {Function} onError - The function to execute if key's value is invalid.
 * @returns {bool} - Returns true if key and key's value is valid.
 */
const validateAndParseFormKey = (key, serializedForm, onError) => {
  if (serializedForm[key] === '') {
    onError('Please enter a valid input.');
    return false;
  }

  switch (key) {
    case 'binary':
      if (!isValidBinary(serializedForm[key])) {
        onError(
          'Please enter a valid binary number, and is not over 64 characters in length.'
        );
        return false;
      }
      break;
    case 'decimal':
      if (!isValidDecimal(serializedForm[key])) {
        onError(
          'Please enter a valid decimal number, and is not over 999,999,999,999.'
        );
        return false;
      }

      // data must be an integer string not a float string
      serializedForm[key] = parseInt(serializedForm[key]);
      break;
    case 'hexadecimal':
      // hexadecimal to uppercase before it reaches the server
      serializedForm[key] = serializedForm[key].toUpperCase();
      if (!isValidHexadecimal(serializedForm[key])) {
        onError(
          'Please enter a valid hexadecimal number, and is not over 64 characters in length.'
        );
        return false;
      }
      break;
    default:
      if (!isValidDecimal(serializedForm[key])) {
        onError('Please enter a valid input, and is not over 999,999,999,999.');
        return false;
      }

      // percentages and total-surface-area category require a float input
      const currentPageURL = window.location.pathname;
      if (
        currentPageURL.indexOf('percentages') !== -1 ||
        currentPageURL.indexOf('total-surface-area') !== -1
      ) {
        serializedForm[key] = parseFloat(serializedForm[key]);
      } else {
        // input is for a int route such as numbers category
        serializedForm[key] = parseInt(serializedForm[key]);
      }
  }
  return true;
};
