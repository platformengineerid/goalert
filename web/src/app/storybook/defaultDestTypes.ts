import { DestinationTypeInfo } from '../../schema'

export const destTypes: DestinationTypeInfo[] = [
  {
    type: 'single-field',
    name: 'Single Field Destination Type',
    enabled: true,
    disabledMessage: 'Single field destination type must be configured.',
    userDisclaimer: '',
    isContactMethod: true,
    isEPTarget: false,
    isSchedOnCallNotify: false,
    iconURL: '',
    iconAltText: '',
    requiredFields: [
      {
        fieldID: 'phone-number',
        labelSingular: 'Phone Number',
        labelPlural: 'Phone Numbers',
        hint: 'Include country code e.g. +1 (USA), +91 (India), +44 (UK)',
        hintURL: '',
        placeholderText: '11235550123',
        prefix: '+',
        inputType: 'tel',
        isSearchSelectable: false,
        supportsValidation: true,
      },
    ],
  },
  {
    type: 'triple-field',
    name: 'Multi Field Destination Type',
    enabled: true,
    disabledMessage: 'Multi field destination type must be configured.',
    userDisclaimer: '',
    isContactMethod: true,
    isEPTarget: false,
    isSchedOnCallNotify: false,
    iconURL: '',
    iconAltText: '',
    requiredFields: [
      {
        fieldID: 'first-field',
        labelSingular: 'First Item',
        labelPlural: 'First Items',
        hint: 'Some hint text',
        hintURL: '',
        placeholderText: '11235550123',
        prefix: '+',
        inputType: 'tel',
        isSearchSelectable: false,
        supportsValidation: true,
      },
      {
        fieldID: 'second-field',
        labelSingular: 'Second Item',
        labelPlural: 'Second Items',
        hint: '',
        hintURL: '',
        placeholderText: 'foobar@example.com',
        prefix: '',
        inputType: 'email',
        isSearchSelectable: false,
        supportsValidation: true,
      },
      {
        fieldID: 'third-field',
        labelSingular: 'Third Item',
        labelPlural: 'Third Items',
        hint: '',
        hintURL: '',
        placeholderText: 'slack user ID',
        prefix: '',
        inputType: 'string',
        isSearchSelectable: false,
        supportsValidation: true,
      },
    ],
  },
  {
    type: 'disabled-destination',
    name: 'This is disabled',
    enabled: false,
    disabledMessage: 'This field is disabled.',
    userDisclaimer: '',
    isContactMethod: true,
    isEPTarget: true,
    isSchedOnCallNotify: true,
    iconURL: '',
    iconAltText: '',
    requiredFields: [
      {
        fieldID: 'disabled',
        labelSingular: '',
        labelPlural: '',
        hint: '',
        hintURL: '',
        placeholderText: 'This field is disabled.',
        prefix: '',
        inputType: 'url',
        isSearchSelectable: false,
        supportsValidation: false,
      },
    ],
  },
]
