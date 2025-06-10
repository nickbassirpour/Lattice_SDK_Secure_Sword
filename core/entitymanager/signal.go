package entitymanager

import (
	components "github.com/nickbassirpour/Lattice_SDK_Secure_Sword.git/api/entities_grpc"
)

func UpdateSignal(entity *components.Entity, new_signal *components.Signal) error {
	if new_signal.FrequencyCenter != nil {
		if new_signal.FrequencyCenter.FrequencyHz != nil {
			new_frequency_hz := new_signal.FrequencyCenter.FrequencyHz
			existing_frequency_hz := entity.Signal.FrequencyCenter.FrequencyHz
			err := UpdateFrequencyHz(existing_frequency_hz, new_frequency_hz)
			if err != nil {
				return err
			}
		}
	}
	if new_signal.FrequencyRange != nil {
		if new_signal.FrequencyRange.MaximumFrequencyRangeHz != nil {
			new_max_freq_range := new_signal.FrequencyRange.MaximumFrequencyRangeHz
			existing_max_freq_range := entity.Signal.FrequencyRange.MaximumFrequencyRangeHz
			if new_max_freq_range.FrequencyHz != nil {
				err := UpdateFrequencyHz(existing_max_freq_range.FrequencyHz, new_max_freq_range.FrequencyHz)
				if err != nil {
					return err
				}
			}
		}
		if new_signal.FrequencyRange.MinimumFrequencyRangeHz != nil {
			new_min_freq_range := new_signal.FrequencyRange.MinimumFrequencyRangeHz
			existing_min_freq_range := entity.Signal.FrequencyRange.MinimumFrequencyRangeHz
			if new_min_freq_range.FrequencyHz != nil {
				err := UpdateFrequencyHz(existing_min_freq_range.FrequencyHz, new_min_freq_range.FrequencyHz)
				if err != nil {
					return err
				}
			}
		}
	}
	if new_signal.BandwidthHz != nil {
		entity.Signal.BandwidthHz = new_signal.BandwidthHz
	}
	if new_signal.SignalToNoiseRatio != nil {
		entity.Signal.SignalToNoiseRatio = new_signal.SignalToNoiseRatio
	}
	if new_signal.LineOfBearing != nil {
		err := UpdateLineOfBearing(entity.Signal.LineOfBearing, new_signal.LineOfBearing)
		if err != nil {
			return nil
		}
	}
	if new_signal.Fixed != nil {
		entity.Signal.Fixed = new_signal.Fixed
	}
	if new_signal.EmitterNotations != nil && len(new_signal.EmitterNotations) > 0 {
		entity.Signal.EmitterNotations = append(entity.Signal.EmitterNotations, new_signal.EmitterNotations...)
	}
	if new_signal.PulseWidthS != nil {
		entity.Signal.PulseWidthS = new_signal.PulseWidthS
	}
	if new_signal.PulseRepetitionInterval != nil && new_signal.PulseRepetitionInterval.PulseRepetitionIntervalS != nil {
		err := UpdatePulseRepetitionInterval(entity.Signal.PulseRepetitionInterval, new_signal.PulseRepetitionInterval)
		if err != nil {
			return nil
		}
	}
	if new_signal.ScanCharacteristics != nil {
		err := UpdateScanCharacteristics(entity.Signal.ScanCharacteristics, new_signal.ScanCharacteristics)
		if err != nil {
			return nil
		}
	}
	return nil
}

func UpdateFrequencyHz(existing_freq_hz *components.FrequencyHz, new_freq_hz *components.FrequencyHz) error {
	if new_freq_hz.Sigma != nil {
		existing_freq_hz.Sigma = new_freq_hz.Sigma
	}
	if new_freq_hz.Value != nil {
		existing_freq_hz.Value = new_freq_hz.Value
	}
	return nil
}

func UpdatePulseRepetitionInterval(existing_pulse_repe_int *components.PulseRepetitionInterval, new_pulse_repe_int *components.PulseRepetitionInterval) error {
	if new_pulse_repe_int.PulseRepetitionIntervalS.Sigma != nil {
		existing_pulse_repe_int.PulseRepetitionIntervalS.Sigma = new_pulse_repe_int.PulseRepetitionIntervalS.Sigma
	}
	if new_pulse_repe_int.PulseRepetitionIntervalS.Value != nil {
		existing_pulse_repe_int.PulseRepetitionIntervalS.Value = new_pulse_repe_int.PulseRepetitionIntervalS.Value
	}
	return nil
}

func UpdateScanCharacteristics(existing_scan_char *components.ScanCharacteristics, new_scan_char *components.ScanCharacteristics) error {
	if new_scan_char.ScanType != nil {
		existing_scan_char.ScanType = new_scan_char.ScanType
	}
	if new_scan_char.ScanPeriodS != nil {
		existing_scan_char.ScanPeriodS = new_scan_char.ScanPeriodS
	}
	return nil
}
