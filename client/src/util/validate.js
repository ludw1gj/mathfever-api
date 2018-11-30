/**
 * Check if binary value is valid.
 * @param {string} binary
 * @returns {bool}
 */
export const isValidBinary = binary => {
  if (binary.length > 64) {
    return false;
  }

  for (let i = 0; i < binary.length; i++) {
    if (binary[i] !== '0' && binary[i] !== '1') {
      return false;
    }
  }
  return true;
};

/**
 * Check if decimal value is valid.
 * @param {string} decimal
 * @returns {bool}
 */
export const isValidDecimal = decimal => {
  return !isNaN(decimal) && decimal < 999999999999;
};

/**
 * Check if hexadecimal value is valid.
 * @param {string} hexadecimal
 * @returns {bool}
 */
export const isValidHexadecimal = hexadecimal => {
  return /^[A-F0-9]+$/.test(hexadecimal) && hexadecimal.length < 64;
};
